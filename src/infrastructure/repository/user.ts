import type { Pool } from "pg";
import { type GoogleId, UserId } from "@/domain/user.ts";
import type { UserRepository as IUserRepository } from "@/domain/user-repository.ts";

export class UserRepository implements IUserRepository {
	private readonly db: Pool;

	constructor(db: Pool) {
		this.db = db;
	}

	async create(googleId: GoogleId): Promise<UserId> {
		const query = `
            INSERT INTO users (google_id)
            VALUES ($1) 
            RETURNING id;
        `;

		const result = await this.db.query(query, [String(googleId)]);

		const row = result.rows[0];
		if (!row) {
			throw new Error("Failed to insert user: no id returned");
		}

		const userIdInt = Number(row.id);
		if (!Number.isFinite(userIdInt)) {
			throw new Error(`Invalid id returned from DB: ${row.id}`);
		}

		const userId = new UserId(userIdInt);
		return userId;
	}
}
