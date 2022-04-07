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

<script setup>
import FridgeButton from '../components/atoms/Button.vue'
import axios from 'axios'
import { ref, onMounted, inject } from 'vue'

const searchedFood = ref('')
const searchResults = ref([]);
const yourFridge = ref([]);
const spoonacularKey = inject('myKey');

onMounted(() => {
    //api al db per prendere il frigo dal db
    axios.get('http://localhost:3000/fridge')
        .then(response => {
            yourFridge.value = response.data.foodList[0].food;
        });
})

function searchItem() {
    axios.get(`https://api.spoonacular.com/food/ingredients/search?query=${searchedFood.value}&number=1`, {
        params: {
            'apiKey' : spoonacularKey
        }
        })
        .then(response => {
            searchResults.value = response.data.results;
        });
}

function addToFridge(food) {
    let checkDuplicates = yourFridge.value.filter(foodItem => {
        return foodItem.Id === food.Id;
    });
    if (checkDuplicates.length === 0) {
        food.quantity = 1;
        yourFridge.value.push(food);
    } else {
        let foodToAddIndex = yourFridge.value.indexOf(checkDuplicates[0]);
        yourFridge.value[foodToAddIndex].quantity++;
    }
}

function removeFromFridge(foodId, index) {
    if (yourFridge.value[index].quantity > 1) {
        yourFridge.value[index].quantity--;
    } else {
        yourFridge.value.splice(index, 1); 
    }
}

function saveNewFridgeItems() {
    axios.post(`http://localhost:3000/fridge/upsert/2`, {
        headers: {
            'Content-Type': 'application/json'
        },
        fridge: yourFridge.value,
    })
    .then(response => {
        console.log(response.data);
    });
}
</script>