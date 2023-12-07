<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="scrollable flex flex-col pt-1.5">
	{#if data.collectionData.entries && data.collectionData.entries.length > 0}
		{#each data.collectionData.entries.sort((a, b) => {
			let aSeen = a.seen;
			let bSeen = b.seen;

			if (aSeen && !bSeen) return 1;
			if (bSeen && !aSeen) return -1;
			return 0;
		}) as entry (entry.id)}
			<a
				class="block border-b border-b-gray-300 bg-gray-200 py-4 hover:bg-gray-400 {entry.seen
					? 'bg-gray-500'
					: ''}"
				href="/entries/{entry.id}"
				>{entry.title}
			</a>
		{/each}
	{:else}
		<h1>No entries</h1>
	{/if}
</div>
