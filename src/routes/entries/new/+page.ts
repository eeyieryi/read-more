import type { PageLoad } from './$types';
import type { Collection } from '$lib/server/db';

export const load: PageLoad = async ({ fetch, url }) => {
	const data = await fetch('/api/collections');
	return {
		collections: (await data.json()) as Collection[],
		fromCollectionId: url.searchParams.get('fromCollection')
	};
};
