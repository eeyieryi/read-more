import { test, expect } from '@playwright/test';

test('blank page has textbox', async ({ page }) => {
	await page.goto('/blank');
	await page.getByRole('textbox').click();
	const want = '最近怎么样？';
	await page.getByRole('textbox').fill(want);
	await expect(page.getByRole('textbox')).toHaveValue(want);
});
