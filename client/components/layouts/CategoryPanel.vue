<template>
  <div class="panel-wrapper">
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>Item</v-expansion-panel-header>
          <v-expansion-panel-content>
            <ThreeCard v-for="data in datas" :key="data.id" :data="data"/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
  </div>
  <!-- /.panel-wrapper -->
</template>

<script>
import ThreeCard from "~/components/ui/ThreeCard";

export default {
  props: ["datas"],
  components: {
    ThreeCard
  }
}
</script>

<style lang="scss" scoped>
.panel-wrapper {
  margin: 30px auto;
}
</style>