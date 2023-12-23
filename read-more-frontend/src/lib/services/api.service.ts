import collectionRepo from './collection';
import entryRepo from './entry';
import uploadRepo from './upload';

const apiService = {
	collection: collectionRepo,
	entry: entryRepo,
	upload: uploadRepo
};
export default apiService;
