<script lang="ts">
	import EntryList from './EntryList.svelte';
	import CollectionDetail from './CollectionDetail.svelte';
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_API } from '$env/static/public';

	export let data: PageData;

	const deleteCollection = async () => {
		if (
			!confirm(
				`Do you really want to delete ${data.collectionData.title} collection?`
			)
		) {
			return;
		}

		const res = await fetch(
			`${PUBLIC_BACKEND_API}/collections/${data.collectionData.id}`,
			{
				method: 'DELETE'
			}
		);
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
