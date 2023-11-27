import { PrismaClient } from '@prisma/client';
import type { Prisma } from '@prisma/client';

export type { Collection, Entry } from '@prisma/client';

export type EntryData = Prisma.EntryGetPayload<{
	include: {
		collection: true;
	};
}>;

export type CollectionData = Prisma.CollectionGetPayload<{
	include: {
		entries: true;
	};
}>;

export const prisma = new PrismaClient();

export default prisma;
