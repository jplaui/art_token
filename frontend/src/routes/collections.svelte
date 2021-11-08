<script context="module">
	import { browser, dev } from '$app/env';

	// we don't need any JS on this page, though we'll load
	// it in dev so that we get hot module replacement...
	export const hydrate = dev;

	// ...but if the client-side router is already loaded
	// (i.e. we came here from elsewhere in the app), use it
	export const router = browser;

	// since there's no dynamic data here, we can prerender
	// it so that it gets served as a static asset in prod
	export const prerender = true;
</script>

<svelte:head>
	<title>ART Token</title>
</svelte:head>

<div class="content">
	<h1>Todos</h1>

	<p>
		User management: creators need a login to upload their collection to IPFS
	</p>

	<pre>open</pre>

	<p>
		User management: login and logout for creators
	</p>

	<pre>open</pre>

	<p>
		User management: users click start to activate the smart contract and their minting protocol
	</p>

	<pre>open</pre>

	<p>
		User management: We do user registration upon request, they will be able to login, we need login button on starting page 
	</p>

	<pre>open</pre>

</div>

<style>
	.content {
		width: 100%;
		max-width: var(--column-width);
		margin: var(--column-margin-top) auto 0 auto;
	}
</style>
