import { test, expect } from "@playwright/test";

test.use({
	ignoreHTTPSErrors: true,
});

test("test", async ({ page }) => {
	await page.goto("http://localhost/local");

	// Set green as starting player.
	const gameCookie = (await page.context().cookies()).find(
		(cookie) => cookie.name == "game",
	)!;
	gameCookie.value = JSON.stringify({
		...JSON.parse(decodeURIComponent(gameCookie.value)),
		currentPlayer: 1,
	});
	await page.context().addCookies([gameCookie]);

	// Reload the page after cookie change.
	await page.reload();

	// First green move.
	await page.locator("td:nth-child(5) > button").first().click();
	await page.locator("td:nth-child(6) > button").first().click();

	// Expect no error (the green player moved a pawn successfully).
	await expect(page.locator(".error")).not.toBeVisible();

	// First red move.
	await page.locator(".green-target > button").first().click();
	await page.locator("tr:nth-child(5) > td:nth-child(7) > button").click();

	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(5) > .red-target > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(5) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td > button").first().click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td > button").first().click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td > button").first().click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(5) > .red-target > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("td:nth-child(4) > button").first().click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td > button").first().click();
	await page.locator("td:nth-child(3) > button").first().click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("button").first().click();
	await page.locator("td:nth-child(3) > button").first().click();
	await page.locator("tr:nth-child(8) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("form").filter({ hasText: "End turn" }).click();
	await page.locator("tr:nth-child(8) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(8) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(8) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(7) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(8) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(7) > button").click();
	await page.locator(".green-target > button").first().click();
	await page.locator("tr:nth-child(7) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.locator(".green-target > button").first().click();
	await page.locator("tr:nth-child(7) > td:nth-child(8) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(8) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("td:nth-child(3) > button").first().click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td > button").first().click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td > button").first().click();
	await page.locator("button").first().click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(4) > td:nth-child(8) > button").click();
	await page.locator(".green-target > button").first().click();
	await page.locator("tr:nth-child(6) > td > button").first().click();
	await page.locator("tr:nth-child(4) > td > button").first().click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(7) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(8) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.getByRole("cell", { name: "2" }).click();
	await page.locator("td:nth-child(3) > button").first().click();
	await page.locator("tr:nth-child(4) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(7) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(5) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("td:nth-child(5) > button").first().click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(5) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(3) > button").click();
	await page.locator("tr:nth-child(6) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(3) > button").click();
	await page.getByRole("button", { name: "End turn" }).click();
	await page.locator("tr:nth-child(6) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(7) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(5) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("td:nth-child(6) > button").first().click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(5) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("tr:nth-child(3) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(6) > button").click();
	await page.locator("tr:nth-child(4) > td:nth-child(4) > button").click();
	await page.locator("td:nth-child(4) > button").first().click();

	// Still has no winner.
	await expect(page.locator(".winner")).not.toBeVisible();

	await page.getByRole("button", { name: "End turn" }).click();

	// Red is the winner.
	await expect(page.locator(".winner.red")).toBeVisible();

	// Cannot move anymore.
	await expect(page.locator('button[value="c5"]')).toBeDisabled();
	await expect(page.locator('button[value="b5"]')).toBeDisabled();
});
