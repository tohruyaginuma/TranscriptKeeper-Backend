export const DB_CONFIG = {
	host: process.env.DB_HOST || "localhost",
	port: parseInt(process.env.DB_PORT || "5432", 10),
	user: process.env.DB_USER || "transcript_keeper_local",
	password: process.env.DB_PASSWORD || "password",
	database: process.env.DB_DATABASE || "transcript_keeper_local",
};
export const PORT = process.env.APP_PORT || 3000;
