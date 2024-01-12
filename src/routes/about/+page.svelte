<script lang="ts">
import Navbar from "$lib/navbar.svelte"
import {Card} from "flowbite-svelte"
import platina from "$lib/assets/PlatinaSB.jpg"
import boy from "$lib/assets/153084275.png"
import cokro from "$lib/assets/133976274.png"

</script>

<Navbar/>
<div class="flex justify-center items-center align-middle">
<Card class="mx-4 flex justify-center items-center font-semibold" href="https://github.com/PlatinaSB">
    
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img src={platina} class="size-32"  alt="platina picture" />
    <h1>Platina Setia Budi</h1>
    <h1>2IA17</h1>
    <h1>51422289</h1>
    <p>CTO & Head Engineer</p>

</Card>

<Card class="mx-4 flex justify-center items-center font-semibold" href="https://github.com/boyishwara">
    
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img src={boy} class="size-32"  alt="platina picture" />
    <h1>Boy Ishwara Aditama</h1>
    <h1>2IA17</h1>
    <h1>50422339</h1>
    <p>CEO & Head Creative</p>

</Card>

<Card class="mx-4 flex justify-center items-center font-semibold" href="https://github.com/Krieto">
    
    <!-- svelte-ignore a11y-img-redundant-alt -->
    <img src={cokro} class="size-32"  alt="platina picture" />
    <h1>Cokro Bayu Laksono</h1>
    <h1>2IA17</h1>
    <h1>50422364</h1>
    <p>Product Manager</p>

</Card></div>