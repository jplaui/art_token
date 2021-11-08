<script context="module">
	export const prerender = true;
</script>

<script>
	// import Counter from '$lib/Counter.svelte';
</script>

<svelte:head>
	<title>ART Token</title>
</svelte:head>

	<!-- Icon Bar (Sidebar - hidden on small screens) -->
	<nav class="w3-sidebar w3-bar-block w3-small w3-hide-small w3-center">
	  <!-- Avatar image in top left corner -->
	  <!-- <img src="/w3images/avatar_smoke.jpg" style="width:100%"> -->
	  <a href="#" class="w3-bar-item w3-button w3-padding-large w3-black">
	    <i class="fa fa-home w3-xxlarge"></i>
	    <p>HOME</p>
	  </a>
	  <a href="#collections" class="w3-bar-item w3-button w3-padding-large w3-hover-black">
	    <i class="fa fa-eye w3-xxlarge"></i>
	    <p>COLLECTIONS</p>
	  </a>
	  <a href="#about" class="w3-bar-item w3-button w3-padding-large w3-hover-black">
	    <i class="fa fa-magic w3-xxlarge"></i>
	    <p>CREATORS</p>
	  </a>
	  <a href="#team" class="w3-bar-item w3-button w3-padding-large w3-hover-black">
	    <i class="fa fa-users w3-xxlarge"></i>
	    <p>TEAM</p>
	  </a>
	  <a href="#about" class="w3-bar-item w3-button w3-padding-large w3-hover-black">
	    <i class="fa fa-question-circle w3-xxlarge"></i>
	    <p>ABOUT</p>
	  </a>
	  <a href="#contact" class="w3-bar-item w3-button w3-padding-large w3-hover-black">
	    <i class="fa fa-envelope w3-xxlarge"></i>
	    <p>CONTACT</p>
	  </a>
	</nav>

	<!-- Navbar on small screens (Hidden on medium and large screens) -->
	<div class="w3-top w3-hide-large w3-hide-medium" id="myNavbar">
	  <div class="w3-bar w3-black w3-opacity w3-hover-opacity-off w3-center w3-small">
	    <a href="#" class="w3-bar-item w3-button" style="width:25% !important">HOME</a>
	    <a href="#collections" class="w3-bar-item w3-button" style="width:25% !important">COLLECTIONS</a>
	    <a href="#creators" class="w3-bar-item w3-button" style="width:25% !important">CREATORS</a>
	    <a href="#team" class="w3-bar-item w3-button" style="width:25% !important">TEAM</a>
	    <a href="#about" class="w3-bar-item w3-button" style="width:25% !important">ABOUT</a>
	    <a href="#contact" class="w3-bar-item w3-button" style="width:25% !important">CONTACT</a>
	  </div>
	</div>

	<!-- Page Content -->
	<div class="w3-padding-large" id="main">
	  <!-- Header/Home -->
	  <header class="w3-container w3-padding-32 w3-center w3-black" id="home">
	    <h1 class="w3-jumbo"><span class="w3-hide-small">Demo:</span> ART Token</h1>
	    <p>AMM supported minting.</p>
	    <img src="/trump.jpg" alt="boy" class="w3-image" width="300" height="500">
	  </header>

	  <!-- Collections Section -->
	  <div class="w3-content w3-justify w3-text-grey w3-padding-64" id="about">
	    <h2 class="w3-text-light-grey">Collections</h2>
	    <hr style="width:200px" class="w3-opacity">
	    <p>Jump over to our collections page and see all active minting protocols where you can participate and take your chance in winning a NFT. Connect your wallet, participate in the random selection protocol to win a NFT and increase your winning chances by staking a specific collection.
	    </p>
	    <div class="w3-row-padding" style="margin:0 -16px">
	      <div class="w3-half">
	        <img src="/blocks.jpg" style="width:100%">
	        <img src="/vr.jpg" style="width:100%">
	        <img src="/world.jpg" style="width:100%">
	      </div>

	      <div class="w3-half">
	        <img src="/art1.jpg" style="width:100%">
	        <img src="/pixel.jpg" style="width:100%">
	        <!-- <img src="/btc.jpg" style="width:100%"> -->
	        <img src="/art2.jpg" style="width:100%">
	      </div>
	    </div>

	    <h3 class="w3-padding-16 w3-text-light-grey">Start Minting</h3>

	    <button class="w3-button w3-light-grey w3-padding-large w3-section">
	      <i class="fa fa-download"></i> Browse Collections
	    </button>

	  <!-- End Collections Section -->
	  </div>


	  <!-- Creators Section -->
	  <div class="w3-content w3-justify w3-text-grey w3-padding-64" id="creators">
	    <h2 class="w3-text-light-grey">Creators</h2>
	    <hr style="width:200px" class="w3-opacity">
	    <p>Creators can register to create and manage their collections. Once a collections is ready, they can deploy it to the contract and start the minting procedure. When minting NFTs of a collections, the creators earn an amount of money which allows them to produce more collections. After creators have registered, they can use the login button in the left navigation bar to manage their collections.
	    </p>

	    <!-- Grid for pricing tables -->
	    <h3 class="w3-padding-16 w3-text-light-grey">Earnings Model</h3>
	    <div class="w3-row-padding" style="margin:0 -16px">
	      <div class="w3-half w3-margin-bottom">
	        <ul class="w3-ul w3-white w3-center w3-opacity w3-hover-opacity-off">
	          <li class="w3-dark-grey w3-xlarge w3-padding-32">Basic</li>
	          <li class="w3-padding-16">Single Collection</li>
	          <li class="w3-padding-16">Images</li>
	          <li class="w3-padding-16">5GB Storage</li>
	          <li class="w3-padding-16">Mail Support</li>
	          <li class="w3-padding-16">
	            <h2>$ 10</h2>
	            <span class="w3-opacity">per month</span>
	          </li>
	          <li class="w3-light-grey w3-padding-24">
	            <button class="w3-button w3-white w3-padding-large w3-hover-black">Sign Up</button>
	          </li>
	        </ul>
	      </div>

	      <div class="w3-half">
	        <ul class="w3-ul w3-white w3-center w3-opacity w3-hover-opacity-off">
	          <li class="w3-dark-grey w3-xlarge w3-padding-32">Pro</li>
	          <li class="w3-padding-16">Multiple Collection</li>
	          <li class="w3-padding-16">Images + 3D Objects</li>
	          <li class="w3-padding-16">50GB Storage</li>
	          <li class="w3-padding-16">Endless Support</li>
	          <li class="w3-padding-16">
	            <h2>$ 25</h2>
	            <span class="w3-opacity">per month</span>
	          </li>
	          <li class="w3-light-grey w3-padding-24">
	            <button class="w3-button w3-white w3-padding-large w3-hover-black">Sign Up</button>
	          </li>
	        </ul>
	      </div>
	    <!-- End Grid/Pricing tables -->
	    </div>

	    <button class="w3-button w3-light-grey w3-padding-large w3-section">
	      <i class="fa fa-download"></i> Register Here
	    </button>

	  <!-- End Creators Section -->
	  </div>


	  <!-- Team Section -->
	  <div class="w3-padding-64 w3-content" id="team">
	    <h2 class="w3-text-light-grey">Team</h2>
	    <hr style="width:200px" class="w3-opacity">

	    <!-- Testimonials -->
	    <!-- <h3 class="w3-padding-24 w3-text-light-grey">My Reputation</h3>   -->
	    <img src="https://www.animaatjes.de/disney-bilder/disney-bilder/donald-duck/donald24.gif" alt="Avatar" class="w3-left w3-circle w3-margin-right" style="width:80px">
	    <p><span class="w3-large w3-margin-right">Kaihua Qin.</span> Imperial College London.</p>
	    <p>DeFi Mastermind.</p><br>
	    
	    <img src="https://www.animaatjes.de/disney-bilder/disney-bilder/donald-duck/dis19.gif" alt="Avatar" class="w3-left w3-circle w3-margin-right" style="width:80px">
	    <p><span class="w3-large w3-margin-right">Liyi Zhou.</span> Imperial College London.</p>
	    <p>DeFi Genius.</p><br>

	    <img src="https://www.animaatjes.de/disney-bilder/disney-bilder/tick-trick-und-track/kkk23.gif" alt="Avatar" class="w3-left w3-circle w3-margin-right" style="width:80px">
	    <p><span class="w3-large w3-margin-right">Jens Ernstberger.</span> Technical University Munich.</p>
	    <p>Managing Director.</p><br>

	    <img src="https://www.animaatjes.de/disney-bilder/disney-bilder/donald-duck/dis20.gif" alt="Avatar" class="w3-left w3-circle w3-margin-right" style="width:80px">
	    <p><span class="w3-large w3-margin-right">Jan Lauinger.</span> Technical University Munich.</p>
	    <p>Full Stack.</p><br>

	    <img src="https://www.animaatjes.de/disney-bilder/disney-bilder/dagobert-duck/dagobert02.gif" alt="Avatar" class="w3-left w3-circle w3-margin-right" style="width:80px">
	    <p><span class="w3-large w3-margin-right">Arthur Gervais.</span> Imperial College London.</p>
	    <p>Angel Whale.</p><br>

	  <!-- End Team Section -->
	  </div>

	  <!-- About Section -->
	  <div class="w3-content w3-justify w3-text-grey w3-padding-64" id="about">
	    <h2 class="w3-text-light-grey">About</h2>
	    <hr style="width:200px" class="w3-opacity">
	    <p>Our goal is to create a decentralized marketplace for non-fungible tokens by simplifying the interaction between liquidity providers and art creators. The aim is to ensure transparency of token Mints, by reducing the unfair distribution of NFTs due to network layer unfairness as well as (insider) information asymmetry during recent NFT mints where automated bots and internal whitelists contributed to significant unfairness.
	    </p>
	    <h3 class="w3-padding-16 w3-text-light-grey">Statistics</h3>
	    <p class="w3-wide">Minted Collections</p>
	    <div class="w3-white">
	      <div class="w3-dark-grey" style="height:28px;width:95%"></div>
	    </div>
	    <p class="w3-wide">Registered Artists</p>
	    <div class="w3-white">
	      <div class="w3-dark-grey" style="height:28px;width:85%"></div>
	    </div>
	    <p class="w3-wide">Number Different Winners</p>
	    <div class="w3-white">
	      <div class="w3-dark-grey" style="height:28px;width:80%"></div>
	    </div><br>
	    
	    <div class="w3-row w3-center w3-padding-16 w3-section w3-light-grey">
	      <div class="w3-quarter w3-section">
	        <span class="w3-xlarge">11+</span><br>
	        ART Traded
	      </div>
	      <div class="w3-quarter w3-section">
	        <span class="w3-xlarge">55+</span><br>
	        Total NFTs Minted
	      </div>
	      <div class="w3-quarter w3-section">
	        <span class="w3-xlarge">89+</span><br>
	        Active in Staking
	      </div>
	      <div class="w3-quarter w3-section">
	        <span class="w3-xlarge">150+</span><br>
	        Happy Customers
	      </div>
	    </div>

	    <button class="w3-button w3-light-grey w3-padding-large w3-section">
	      <i class="fa fa-download"></i> Download Resume
	    </button>

	  <!-- End About Section -->
	  </div>

	  <!-- Contact Section -->
	  <div class="w3-padding-64 w3-content w3-text-grey" id="contact">
	    <h2 class="w3-text-light-grey">Contact Us</h2>
	    <hr style="width:200px" class="w3-opacity">

	    <div class="w3-section">
	      <p><i class="fa fa-map-marker fa-fw w3-text-white w3-xxlarge w3-margin-right"></i> Crypto Valley, Monaco</p>
	      <p><i class="fa fa-phone fa-fw w3-text-white w3-xxlarge w3-margin-right"></i> Phone: +00 151515</p>
	      <p><i class="fa fa-envelope fa-fw w3-text-white w3-xxlarge w3-margin-right"> </i> Email: mail@mail.com</p>
	    </div><br>
	    <p>Let's get in touch. Send us a message or ETH:</p>

	    <form action="/action_page.php" target="_blank">
	      <p><input class="w3-input w3-padding-16" type="text" placeholder="Name" required name="Name"></p>
	      <p><input class="w3-input w3-padding-16" type="text" placeholder="Email" required name="Email"></p>
	      <p><input class="w3-input w3-padding-16" type="text" placeholder="Subject" required name="Subject"></p>
	      <p><input class="w3-input w3-padding-16" type="text" placeholder="Message" required name="Message"></p>
	      <p>
	        <button class="w3-button w3-light-grey w3-padding-large" type="submit">
	          <i class="fa fa-paper-plane"></i> SEND MESSAGE
	        </button>
	      </p>
	    </form>
	  <!-- End Contact Section -->
	  </div>
	  
	    <!-- Footer -->
	  <footer class="w3-content w3-padding-64 w3-text-grey w3-xlarge">
	    <i class="fa fa-facebook-official w3-hover-opacity"></i>
	    <i class="fa fa-instagram w3-hover-opacity"></i>
	    <i class="fa fa-snapchat w3-hover-opacity"></i>
	    <i class="fa fa-pinterest-p w3-hover-opacity"></i>
	    <i class="fa fa-twitter w3-hover-opacity"></i>
	    <i class="fa fa-linkedin w3-hover-opacity"></i>
	    <p class="w3-medium">Powered by <a href="https://www.w3schools.com/w3css/default.asp" target="_blank" class="w3-hover-text-green">w3.css</a></p>
	  <!-- End footer -->
	  </footer>

	<!-- END PAGE CONTENT -->
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
