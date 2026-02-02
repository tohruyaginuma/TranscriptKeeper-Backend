import type { UserRepository } from "@/domain/user-repository.ts";
import type { GoogleId, UserId } from "@/domain/user.ts";
import type { UserUseCase as IUserUseCase } from "@/application/user-interface.ts";

export class UserUseCase implements IUserUseCase {
	private readonly repo: UserRepository;

	constructor(repo: UserRepository) {
		this.repo = repo;
	}

	async create(googleId: GoogleId): Promise<UserId> {
		const userId = await this.repo.create(googleId);

		return userId;
	}
}
