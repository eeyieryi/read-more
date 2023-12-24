import { test, expect } from '@playwright/test';
import { getPathname } from './utils';

test('index page redirects to /collections', async ({ page }) => {
	await page.goto('/');
	await page.waitForURL('**/collections');
	const got = getPathname(page.url());
	const want = '/collections';
	expect(got).toBe(want);
});
