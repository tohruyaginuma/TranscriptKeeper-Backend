import express from "express";
import { PORT } from "@/lib/config.ts";
import { setRoute } from "@/route/route.ts";

const newExpressApp = () => {
	const app = express();

	return app;
};

const main = () => {
	const app = newExpressApp();

	setRoute(app);

	app.listen(PORT, () => {
		console.log(`Example app listening on port ${PORT}`);
	});
};

main();
