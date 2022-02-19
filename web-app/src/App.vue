<template>
  <FridgeLayout>
    <router-view/>
  </FridgeLayout>
</template>
<script>
import FridgeLayout from './components/layouts/Auth'

export default {
  components: {
    FridgeLayout
  }
}
</script>
