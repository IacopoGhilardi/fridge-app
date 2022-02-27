<template>
    <div class="fridge">
        <div class="mx-auto p-4 w-max">
            <input type="text" placeholder="Add some ingridient.." class="px-6 py-2 border rounded" v-model="searchedFood">
            <button class="bg-red-500 text-white py-2 px-4 rounded" type="button" @click="searchItem">Search</button>
        </div>

        <div class="flex flex-wrap justify-center">
            <div v-for="(food, index) in searchResults" :key="index" class="border p-8">
                <div class="mb-4">
                    <img :src="`https://spoonacular.com/cdn/ingredients_100x100/${food.image}`" alt="food_image">
                </div>
                <div class="text-center">
                    <p>
                        {{ food.name.toUpperCase() }}
                    </p>
                    <FridgeButton colorClass="bg-green-500" text="Aggiungi" @click="addToFridge(food)"></FridgeButton>
                </div>
            </div>
        </div>

        <div>
            <h1 class="text-3xl mb-4 text-center mt-4">Ecco cosa contiene il tuo frigo</h1>
            <div class="flex justify-center">
                <button class="bg-green-500 text-white py-2 px-4 rounded" @click="saveNewFridgeItems">Salva modifiche</button>
            </div>
            <div class="flex flex-wrap">
                <div v-for="(yourFood, index) in yourFridge" :key="index" class="border p-8 pt-0 m-2">
                    <div class="text-center text-xl my-2">
                        x{{ yourFood.quantity }}
                    </div>
                    <div class="mb-4">
                        <img :src="`https://spoonacular.com/cdn/ingredients_100x100/${yourFood.image}`" alt="your_food_image">
                    </div>
                    <div class="text-center">
                        <p>
                            {{ yourFood.name.toUpperCase() }}
                        </p>
                        <FridgeButton colorClass="bg-red-500" text="Rimuovi" @click="removeFromFridge(yourFood.id, index)"></FridgeButton>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import FridgeButton from '../components/atoms/Button.vue'

export default {
    components: {
        FridgeButton,
    },

    props: {

    },

    data() {
        return {
            searchedFood: '',
            searchResults: [],
            yourFridge: [],
        }
    },

    mounted() {   
        //api al db per prendere il frigo dal db
        this.axios.get('http://localhost:3000/fridge')
            .then(response => {
                console.log('your fridge', response.data.fridgeItems);
                console.log('ciaooo', response);
                this.yourFridge = response.data.fridgeItems[0].food;
            });
    },

    methods: {
        searchItem() {
            this.axios.get(`https://api.spoonacular.com/food/ingredients/search?query=${this.searchedFood}&number=1`, {
            params: {
                'apiKey' : this.myKey
            }
            })
            .then(response => {
                this.searchResults = response.data.results;
            });
        },
        addToFridge(food) {
            let checkDuplicates = this.yourFridge.filter(foodItem => {
                return foodItem.id === food.id;
            });
            if (checkDuplicates.length === 0) {
                food.quantity = 1;
                this.yourFridge.push(food);
            } else {
                let foodToAddIndex = this.yourFridge.indexOf(checkDuplicates[0]);
                this.yourFridge[foodToAddIndex].quantity++;
            }
        },
        removeFromFridge(foodId, index) {
            if (this.yourFridge[index].quantity > 1) {
                this.yourFridge[index].quantity--;
            } else {
                this.yourFridge.splice(index, 1); 
            }
        },
        saveNewFridgeItems() {
            this.axios.post(`http://localhost:3000/fridge/upsert/2`, {
                headers: {
                    'Content-Type': 'application/json'
                },
                fridge: this.yourFridge,
            })
            .then(response => {
                console.log(response);
            });
        }
    },
}
</script>