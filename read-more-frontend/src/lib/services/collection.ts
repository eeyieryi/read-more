import type {
	CreateOneReturnType,
	ReturnTypeBase,
	ReturnTypeBaseWithId,
	SuccessfulWithObject,
	Unsuccessful
} from './utils';
import { PUBLIC_BACKEND_API } from '$env/static/public';
import type { collectionModel } from '$lib/models';

type fetch = typeof fetch;

async function createOne(
	fetch: fetch,
	dto: collectionModel.CreateOneDto
): CreateOneReturnType<collectionModel.Collection, 'collection'> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/collections`, {
			method: 'POST',
			body: JSON.stringify(dto),
			headers: {
				'content-type': 'application/json'
			}
		});
		if (res.ok) {
			const data = (await res.json()) as collectionModel.Collection;
			return {
				success: true,
				collection: data
			};
		}
		return { success: false };
	} catch (err) {
		return { success: false };
	}
}

async function deleteOne(
	fetch: fetch,
	id: string
): Promise<ReturnTypeBaseWithId> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/collections/${id}`, {
			method: 'DELETE'
		});
		return { success: res.ok, id };
	} catch (err) {
		return {
			success: false,
			id
		};
	}
}

async function findAll(fetch: fetch): Promise<
	ReturnTypeBase & {
		collections: collectionModel.Collection[];
	}
> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/collections`);
		const data = (await res.json()) as collectionModel.Collection[];
		return { success: res.ok, collections: data };
	} catch (err) {
		return {
			success: false,
			collections: []
		};
	}
}

async function findOne(
	fetch: fetch,
	id: string
): Promise<
	| SuccessfulWithObject<collectionModel.CollectionData, 'collectionData'>
	| Unsuccessful
> {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_API}/collections/${id}`);
		if (res.ok) {
			const data = (await res.json()) as collectionModel.CollectionData;
			return {
				success: true,
				collectionData: data
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

const collectionRepo = {
	createOne,
	deleteOne,
	findAll,
	findOne
};
export default collectionRepo;
