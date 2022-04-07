<template>
 <div class="flex flex-col space-y-4 mt-4">
     <input v-model="email" class="w-max border px-2" type="text" placeholder="Email">
     <input v-model="password" class="w-max border px-2" type="password" placeholder="Password">
        <button class="w-max bg-blue-500 px-4 py-2 text-white rounded-md" @click="login">Login</button>     
</div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

const email = ref('');
const password = ref('');

function login() {
    console.log('sisisi');

    axios.post('http://localhost:3000/login', {
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