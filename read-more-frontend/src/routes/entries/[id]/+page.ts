import type { PageLoad } from './$types';
import type { collectionModel, entryModel } from '$lib/models';
import { PUBLIC_BACKEND_API } from '$env/static/public';

export const load: PageLoad = async ({ fetch, params: { id } }) => {
	return {
		entryData: (await (
			await fetch(`${PUBLIC_BACKEND_API}/entries/${id}`)
		).json()) as entryModel.EntryData,
		collections: (await (
			await fetch(`${PUBLIC_BACKEND_API}/collections`)
		).json()) as collectionModel.Collection[]
	};
};
