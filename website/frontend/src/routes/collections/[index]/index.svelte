<script context="module">
	export async function load({ page }) {
		const { index } = page.params;
		return {
			props: { index }
		};
	}
</script>

<script>
	import { onMount } from 'svelte'
	export let index;

	let nft_index = 10;

	onMount(async () => {
		console.log("page index:", index);
	})

	function enterAuction(index) {
		console.log("test...")
	}

</script>

<svelte:head>
	<title>ART Token</title>
</svelte:head>

<!-- Page Content -->
	<div class="w3-padding-large" id="main">
	<!-- Header/Home -->
	<header class="w3-container w3-padding-32 w3-center w3-black">
		<h1 class="w3-jumbo"><span class="w3-hide-small">Demo:</span> ART Token</h1>
		<p>NFT Auction.</p>
	  </header>

	  <!-- Collections Section -->
	  <div class="w3-content w3-justify w3-black">
			<p><a href="/collections">Back.</a></p>
		</div>

	  <div class="w3-content w3-justify w3-text-grey w3-padding-64">
		<h2 class="w3-text-light-grey">Collection Name</h2>
		<hr style="width:200px" class="w3-opacity">
		<p>NFT number: {nft_index}
		</p>

	    <div class="w3-display-container w3-container">
			<img src="/art1.jpg" alt="Jeans" style="width:100%">
			<div class="w3-display-topleft w3-text-white" style="padding:24px 48px">
			  <!-- <h1 class="w3-jumbo w3-hide-small">New arrivals</h1> -->
			  <!-- <h1 class="w3-hide-large w3-hide-medium">New arrivals</h1> -->
			  <!-- <h1 class="w3-hide-small">COLLECTION 2016</h1> -->
			  <!-- <p><a href="#jeans" class="w3-button w3-black w3-padding-large w3-large">SHOP NOW</a></p> -->
			</div>
		</div>

		<br>
		
		<div class="w3-center">
			<button><a href="#jeans" class="w3-button w3-black w3-padding-large w3-large">Connect Wallet</a></button>
			<br>
			<br>
			<button class="w3-button w3-black" on:click={() => enterAuction(1)}>Place Bid <i class="fa fa-shopping-cart"></i></button>
			<button class="w3-button w3-black" on:click={() => enterAuction(2)}>Staking <i class="fa fa-shopping-cart"></i></button>
		</div>

	<!-- End Collections Section -->
	</div>
	</div>

<style>
	body, h1,h2,h3,h4,h5,h6 {font-family: "Montserrat", sans-serif}
	.w3-row-padding img {margin-bottom: 12px}
	/* Set the width of the sidebar to 120px */
	.w3-sidebar {width: 120px;background: #222;}
	/* Add a left margin to the "page content" that matches the width of the sidebar (120px) */
	#main {margin-left: 120px}
	/* Remove margins from "page content" on small screens */
	@media only screen and (max-width: 600px) {#main {margin-left: 0}}
</style>
