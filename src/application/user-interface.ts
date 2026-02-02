import type { GoogleId, UserId } from "@/domain/user.ts";

export type UserUseCase = {
	create(googleId: GoogleId): Promise<UserId>;
};
