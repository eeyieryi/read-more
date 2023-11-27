import type { LayoutServerLoad } from './$types';
import type { Collection } from '$lib/server/db';

export const load: LayoutServerLoad = async ({ fetch }) => {
	const res = await fetch('/api/collections');
	return {
		collections: ((await res.json()) as Collection[]) || []
	};
};
