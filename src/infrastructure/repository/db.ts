import { Pool } from "pg";
import { DB_CONFIG } from "@/lib/config.js";

export const initDb = (): Pool => {
	const pool = new Pool(DB_CONFIG);

	return pool;
};
