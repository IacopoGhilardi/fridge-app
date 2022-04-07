<template>
 <div class="flex flex-col space-y-4 mt-4">
     <input v-model="firstName" class="w-max border" type="text" placeholder="First name">
     <input v-model="lastname" class="w-max border" type="text" placeholder="Last name">
     <input v-model="email" class="w-max border" type="text" placeholder="Email">
     <input v-model="password" class="w-max border" type="password" placeholder="Password">

     <button class="w-max bg-blue-500 px-4 py-2 text-white rounded-md" @click="clickButton">Register</button>
 </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const firstName = ref('');
const lastname = ref('');
const email = ref('');
const password = ref('');

function clickButton() {
    console.log('sisisi');
    console.log(firstName.value);

    axios.post('http://localhost:3000/register', {
        first_name: firstName.value,
        last_name: lastname.value,
        email: email.value,
        password: password.value,
    })
    .then(response => {
        console.log('response', response);
    })
    .catch(function (error) {
        console.log(error);
    })

}

</script>