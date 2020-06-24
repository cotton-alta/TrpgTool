<template>
  <div class="panel-wrapper">
        <!-- class="panel-wrapper" -->
      <v-expansion-panels
        accordion
      >
        <v-expansion-panel>
          <v-expansion-panel-header>ダイス</v-expansion-panel-header>
          <v-expansion-panel-content>
            <AbilityCard v-for="data in datas" :key="data.id" :data="data"/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
  </div>
  <!-- /.panel-wrapper -->
</template>

<script>
import AbilityCard from "~/components/ui/AbilityCard";

export default {
  props: ["datas"],
  components: {
    AbilityCard
  }
}
</script>

<style lang="scss" scoped>
.panel-wrapper {
  margin: 30px auto;
}
</style>