<script lang="ts">
	import EntryList from './EntryList.svelte';
	import CollectionDetail from './CollectionDetail.svelte';
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	export let data: PageServerData;

	const deleteCollection = async () => {
		if (
			!confirm(
				`Do you really want to delete ${data.collectionData.title} collection?`
			)
		) {
			return;
		}

		const res = await fetch(`/api/collections/${data.collectionData.id}`, {
			method: 'DELETE'
		});
		if (res.ok) {
			goto('/');
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
