<template>
    <div>
        <button type="button" class="w-full text-white px-4 py-2 rounded" :class="colorClass">{{ text }}</button>
    </div>
</template>
<script>
export default {
    props: {
        text: String,
        colorClass: {
            type: String,
            default: 'bg-green-500'
        }
    }
}
</script>