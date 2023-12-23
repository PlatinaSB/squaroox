<script lang="ts">
import Navbar from "$lib/navbar.svelte"
import {Card} from "flowbite-svelte"
import platina from "$lib/assets/platinasb.png"

</script>

<Navbar/>

<Card class="mx-4 flex justify-center items-center font-semibold">
    
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img src={platina} class="size-32"  alt="platina picture" />
    <h1>PlatinaSB</h1>
    <h1>0IA00</h1>
    <h1>00000000</h1>
    <p>Lead Developer and Enginer</p>

</Card>