<script lang="ts">
	import EntryList from './EntryList.svelte';
	import CollectionDetail from './CollectionDetail.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import apiService from '$lib/services/api.service';

	export let data: PageData;

	const deleteCollection = async () => {
		if (
			!confirm(
				`Do you really want to delete ${data.collectionData.title} collection?`
			)
		) {
			return;
		}

		const res = await apiService.collection.deleteOne(
			fetch,
			data.collectionData.id
		);
		if (res.success) {
			goto('/');
		} else {
			// TODO: handle error
		}
	};
</script>

<svelte:head>
	<title>{data.collectionData.title}</title>
</svelte:head>

<div class="flex w-full flex-col text-center">
	<CollectionDetail data="{data}" on:delete="{deleteCollection}" />
	<EntryList data="{data}" />
</div>
