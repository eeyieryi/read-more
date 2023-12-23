import type {
	CreateOneReturnType,
	ReturnTypeBase,
	SuccessfulWithObject,
	Unsuccessful
} from './utils';
import { PUBLIC_BACKEND_API } from '$env/static/public';
import type { entryModel } from '$lib/models';

type fetch = typeof fetch;

async function createOne(
	fetch: fetch,
	dto: entryModel.CreateOneDto
): CreateOneReturnType<entryModel.EntryData, 'entryData'> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/entries`, {
			method: 'POST',
			body: JSON.stringify(dto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			const data = (await res.json()) as entryModel.EntryData;
			return {
				success: true,
				entryData: data
			};
		}
		return { success: false };
	} catch (err) {
		return {
			success: false
		};
	}
}

async function deleteOne(fetch: fetch, id: string): Promise<ReturnTypeBase> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/entries/${id}`, {
			method: 'DELETE'
		});
		return { success: res.ok };
	} catch (err) {
		return { success: false };
	}
}

async function findOne(
	fetch: fetch,
	id: string
): Promise<
	SuccessfulWithObject<entryModel.EntryData, 'entryData'> | Unsuccessful
> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/entries/${id}`);
		if (res.ok) {
			const data = (await res.json()) as entryModel.EntryData;
			return {
				success: true,
				entryData: data
			};
		}
		return {
			success: false
		};
	} catch (err) {
		return {
			success: false
		};
	}
}

async function updateOne(
	fetch: fetch,
	dto: entryModel.UpdateOneDto
): Promise<SuccessfulWithObject<entryModel.Entry, 'entry'> | Unsuccessful> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/entries/${dto.id}`, {
			method: 'PUT',
			body: JSON.stringify(dto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			const data = (await res.json()) as entryModel.Entry;
			return {
				success: true,
				entry: data
			};
		}
		return { success: false };
	} catch (err) {
		return {
			success: false
		};
	}
}
const entryRepo = {
	createOne,
	deleteOne,
	findOne,
	updateOne
};
export default entryRepo;
