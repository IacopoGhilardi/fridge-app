<template>
    <div>
        <div class="mb-4">
            <img :src="`https://spoonacular.com/cdn/ingredients_100x100/${food.image}`" alt="food_image">
        </div>
        <div class="text-center">
            <p>
                {{ food.name.toUpperCase() }}
            </p>
            <FridgeButton v-if="type=='add'" colorClass="bg-green-500" text="Aggiungi" @click="addToFridge(food)"></FridgeButton>
            <FridgeButton v-if="type=='remove'" colorClass="bg-green-500" text="Aggiungi" @click="addToFridge(food)"></FridgeButton>
        </div>
    </div>
</template>
<script>
import FridgeButton from '../atoms/Button.vue'
export default {
    components: {
        FridgeButton
    },

    props: {
        type: String,
        food: Object,
    },

    data() {
        return {
            
        }
    },

    mounted() {
    },

    methods: {
        buttonClick() {
            this.$emit();
        }
    },
    
}
</script>
