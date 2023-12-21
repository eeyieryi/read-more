import { test, expect } from '@playwright/test';

test('index page redirects correctly', async ({ page }) => {
	await page.goto('/');
	await page.waitForURL('**/collections');
	const got = new URL(page.url()).pathname;
	const want = '/collections';
	expect(got).toBe(want);
});
