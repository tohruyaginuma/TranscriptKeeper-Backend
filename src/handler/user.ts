import type { UserUseCase } from "@/application/user-interface.ts";

export class UserHandler {
	private readonly service: UserUseCase;

	constructor(service: UserUseCase) {
		this.service = service;
	}

	create() {}
}
