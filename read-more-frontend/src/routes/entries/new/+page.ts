import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import apiService from '$lib/services/api.service';
import { unableToLoadError } from '$lib/errors';

export const load: PageLoad = async ({ fetch, url }) => {
	const res = await apiService.collection.findAll(fetch);
	if (!res.success) {
		const err = unableToLoadError;
		error(err.status, err.message);
	}
	return {
		collections: res.data,
		fromCollectionId: url.searchParams.get('fromCollection')
	};
};
