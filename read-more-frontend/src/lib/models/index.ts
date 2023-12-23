import collectionModel from './collection';
export type * as collectionModel from './collection';

import entryModel from './entry';
export type * as entryModel from './entry';

const models = {
	collection: collectionModel,
	entry: entryModel
};
export default models;
