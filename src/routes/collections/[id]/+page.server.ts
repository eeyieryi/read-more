import type { CollectionData } from '$lib/server/db';

export const load = async ({ fetch, params: { id } }) => {
	const res = await fetch(`/api/collections/${id}`);
	return {
		collectionData: (await res.json()) as CollectionData
	};
};
