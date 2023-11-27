import type { PageServerLoad } from './$types';
import type { Collection, EntryData } from '$lib/server/db';

export const load: PageServerLoad = async ({ fetch, params: { id } }) => {
	return {
		entryData: (await (
			await fetch(`/api/entries/${id}`)
		).json()) as EntryData,
		collections: (await (
			await fetch('/api/collections')
		).json()) as Collection[]
	};
};
