import { PUBLIC_BACKEND_API } from '$env/static/public';
import type { collectionModel } from '$lib/models';

export const load = async ({ fetch, params: { id } }) => {
	const res = await fetch(`${PUBLIC_BACKEND_API}/collections/${id}`);
	return {
		collectionData: (await res.json()) as collectionModel.CollectionData
	};
};
