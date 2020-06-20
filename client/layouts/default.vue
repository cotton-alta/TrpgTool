<template>
  <div>
    <v-app>
      <v-app-bar app

      >
        <v-app-bar-nav-icon @click.stop="drawer = !drawer">
        </v-app-bar-nav-icon>
      </v-app-bar>
      <nuxt class="nuxt-wrapper" />
    </v-app>
  </div>
</template>

<script>
export default {
  data () {
    return {
      drawer: false
    }
  }
}
</script>

<style lang="scss" scoped>
.nuxt-wrapper {
  padding: 100px 0px;
}
</style>