<template lang="">
    <div class="mx-auto min-h-screen">
        <Header></Header>

        <div class="w-full md:w-3/5 mx-auto">
            <slot></slot>
        </div>
    </div>
</template>
<script>
import Header from './partials/Header';

export default {
    components: {
        Header
    },

    props: {

    },

    data() {
        return {
            
        }
    },

    mounted() {
        this.axios.get('https://api.spoonacular.com/food/ingredients/search?query=banana', {
            params: {
                'apiKey' : '49b11a04577b4b818c848eef087c4315'
            }
        })
        .then(response => {
            console.log(response.data);
        });
    },
}
</script>