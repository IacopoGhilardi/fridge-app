<template>
  <div class="home h-full">
    <h1>Ricette</h1>


  </div>
</template>

<script>

export default {
  components: {

  },
  data() {
    return {
      yourFridge: []
    }
  },
  mounted() {
    this.axios.get('http://localhost:3000/fridge')
    .then(response => {
        this.yourFridge = response.data.fridgeItems[0].food;
        console.log('fridge!!', this.yourFridge);
        // this.axios.get('https://api.spoonacular.com/recipes/findByIngredients', {
        //   params: {
        //     'apiKey' : this.myKey,
        //     'ingredients': 'apples,flour,sugar'
        //   }
        // })
        // .then(response => {
        //   console.log('response', response);
        // }); 
    });
  },
}
</script>
