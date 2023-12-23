import { error } from '@sveltejs/kit';
import { unableToLoadError } from '$lib/errors.js';
import apiService from '$lib/services/api.service.js';

export const load = async ({ fetch, params: { id } }) => {
	const res = await apiService.collection.findOne(fetch, id);
	if (!res.success) {
		const err = unableToLoadError;
		error(err.status, err.message);
	}
	return {
		collectionData: res.collectionData
	};
};
