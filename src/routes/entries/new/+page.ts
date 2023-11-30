import type { PageLoad } from './$types';
import { PUBLIC_BACKEND_API } from '$env/static/public';
import type { collectionModel } from '$lib/models';

export const load: PageLoad = async ({ fetch, url }) => {
	const data = await fetch(`${PUBLIC_BACKEND_API}/collections`);
	return {
		collections: (await data.json()) as collectionModel.Collection[],
		fromCollectionId: url.searchParams.get('fromCollection')
	};
};
