<template>
    <div class="flex py-4 px-6 bg-black">
        <div class="none md:block w-4/12"></div>
        <div class="w-6/12 md:w-4/12 flex">
            <input type="text" placeholder="Type an ingidient.." class="border-t border-b border-l px-2 justify-self-center w-full" v-model="searchRecipe">
            <button class="border px-4 text-yellow-400"><i class="fas fa-search"></i></button>
        </div>
        <ul class="justify-end none md:flex w-4/12">
            <li v-for="(element, index) in navbar" :key="index" class="mx-3 text-yellow-400"><router-link :to="element.link">{{ element.title }}</router-link></li>
        </ul>
    </div>
</template>
<script>
export default {
    components: {

    },

    props: {

    },

    data() {
        return {
            navbar: [
                {
                    title: 'Home',
                    link: '/'
                },
                {
                    title: 'Piu viste',
                    link: '#'
                },
                {
                    title: 'Il tuo frigo',
                    link: 'your-fridge'
                },
                {
                    title: 'Profilo',
                    link: '#'
                },
            ],
        }
    },

    mounted() {
    },
}
</script>