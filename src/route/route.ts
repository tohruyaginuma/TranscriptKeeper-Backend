import type { Express } from "express";

export const setRoute = (app: Express) => {
	app.get("/", (req, res) => {
		res.send("Hello World!");
	});
};
