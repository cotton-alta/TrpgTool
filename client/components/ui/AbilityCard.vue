<template>
<div class="card-wrapper">
  <v-container>
    <v-row>
      <v-col>
        <v-card
          class="card-title"
          outlined
          tile
        >
          {{ data.name }}
        </v-card>
        <v-row no-gutters>
          <v-col>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <input type="text" v-model="data.value" />
            </v-card>
          </v-col>
          <v-col>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <input type="text" />
            </v-card>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</div>
<!-- /.card-wrapper -->
</template>

<script>
export default {
  props: ["data"]
}
</script>

<style lang="scss" scoped>
.card {
  &-wrapper {
    height: 80px;
    margin: 20px;
  }
  &-title {
    line-height: 40px;
    height: 40px;
    text-align: center;
  }
  &-content {
    input {
      height: 100%;
      width: 100%;
      text-align: center;
      border: none;
    }
    line-height: 40px;
    height: 40px;
  }
}

v-row {
  height: 100%;
}
v-col {
  height: 100%;
}
</style>