import type { GoogleId, UserId } from "@/domain/user.ts";

export type UserRepository = {
	create(googleId: GoogleId): Promise<UserId>;
};
