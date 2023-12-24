import { test, expect } from '@playwright/test';
import { getPathname } from './utils';

test.describe('root layout', () => {
	test.beforeEach(async ({ page }, testInfo) => {
		console.log(testInfo.title);
		await page.goto('/');
	});
	test.describe('navbar', () => {
		test('has navbar', async ({ page }) => {
			await expect(page.getByRole('navigation')).toBeVisible();
		});
		test('toggles navbar visibility', async ({ page }) => {
			await expect(page.getByRole('navigation')).toBeVisible();

			const testShowing = async () => {
				await expect(page.getByRole('navigation')).toBeVisible();
				await expect(page.getByRole('navigation')).toContainText(
					'Collections Blank New Collection New Entry -'
				);
				await expect(page.getByRole('navigation')).not.toContainText(
					'+'
				);
				await expect(
					page.getByRole('link', { name: 'Collections' })
				).toBeVisible();
				await expect(
					page.getByRole('link', { name: 'Blank' })
				).toBeVisible();
				await expect(
					page.getByRole('link', { name: 'New Collection' })
				).toBeVisible();
				await expect(
					page.getByRole('link', { name: 'New Entry' })
				).toBeVisible();
				await expect(
					page.getByRole('button', { name: '-' })
				).toBeVisible();
				await expect(
					page.getByRole('button', { name: '+' })
				).not.toBeVisible();
			};

			const testHidden = async () => {
				await expect(page.getByRole('navigation')).toBeVisible();
				await expect(page.getByRole('navigation')).not.toContainText(
					'Collections Blank New Collection New Entry -'
				);
				await expect(page.getByRole('navigation')).toContainText('+');
				await expect(
					page.getByRole('link', { name: 'Collections' })
				).not.toBeVisible();
				await expect(
					page.getByRole('link', { name: 'Blank' })
				).not.toBeVisible();
				await expect(
					page.getByRole('link', { name: 'New Collection' })
				).not.toBeVisible();
				await expect(
					page.getByRole('link', { name: 'New Entry' })
				).not.toBeVisible();
				await expect(
					page.getByRole('button', { name: '-' })
				).not.toBeVisible();
				await expect(
					page.getByRole('button', { name: '+' })
				).toBeVisible();
			};

			// navbar is not hidden by default
			await testShowing();

			for (let i = 0; i < 10; i++) {
				// hide it
				await page.getByRole('button', { name: '-' }).click();
				await testHidden();

				// show it
				await page.getByRole('button', { name: '+' }).click();
				await testShowing();
			}
		});

		test('navigation', async ({ page }) => {
			// root page is /collections
			await page.waitForURL('**/collections');
			expect(getPathname(page.url())).toBe('/collections');
			await expect(
				page.getByRole('link', { name: 'Collections' })
			).toHaveClass(/active/);

			await page.getByRole('link', { name: 'Blank' }).click();
			await page.waitForURL('**/blank');
			expect(getPathname(page.url())).toBe('/blank');
			await expect(page.getByRole('link', { name: 'Blank' })).toHaveClass(
				/active/
			);

			await page.getByRole('link', { name: 'Collections' }).click();
			await page.waitForURL('**/collections');
			expect(getPathname(page.url())).toBe('/collections');
			await expect(
				page.getByRole('link', { name: 'Collections' })
			).toHaveClass(/active/);

			await page.getByRole('link', { name: 'New Collection' }).click();
			await page.waitForURL('**/collections/new');
			expect(getPathname(page.url())).toBe('/collections/new');
			await expect(
				page.getByRole('link', { name: 'New Collection' })
			).toHaveClass(/active/);

			await page.getByRole('link', { name: 'New Entry' }).click();
			await page.waitForURL('**/entries/new');
			expect(getPathname(page.url())).toBe('/entries/new');
			await expect(
				page.getByRole('link', { name: 'New Entry' })
			).toHaveClass(/active/);
		});
	});
});
