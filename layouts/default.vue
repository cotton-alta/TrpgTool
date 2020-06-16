<template>
  <div>
    <v-app>
      <v-app-bar app>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer">
        </v-app-bar-nav-icon>
      </v-app-bar>
      <nuxt />
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
