import type { LayoutLoad } from './$types';
import type { collectionModel } from '$lib/models';
import { PUBLIC_BACKEND_API } from '$env/static/public';

export const load: LayoutLoad = async ({ fetch }) => {
	const res = await fetch(`${PUBLIC_BACKEND_API}/collections`);
	return {
		collections: ((await res.json()) as collectionModel.Collection[]) || []
	};
};
