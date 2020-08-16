<template>
<div class="card-wrapper">
  <v-container>
    <v-row>
      <v-col>
        <v-row no-gutters>
          <v-col>
            <v-card
              class="card-content"
              outlined
              tile
            >
              {{ data.tribe.name }}
            </v-card>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <select>
                <option value="只人">只人</option>
              </select>
            </v-card>
          </v-col>
          <v-col>
            <v-card
              class="card-content"
              outlined
              tile
            >
              {{ data.gender.name }}
            </v-card>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <select>
                <option value="男">男</option>
                <option value="女">女</option>
              </select>
            </v-card>
          </v-col>
          <v-col>
            <v-card
              class="card-content"
              outlined
              tile
            >
              {{ data.age.name }}
            </v-card>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <input type="text" v-model="data.age.value" />
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
  props: [
    "data"
  ],
  data() {
    return {
      dropdown_font: ['Arial', 'Calibri', 'Courier', 'Verdana']
    }
  }
}
</script>

<style lang="scss" scoped>
.card {
  &-wrapper {
    margin: 10px;
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
      border: none;
      text-align: center;
    }
    select {
      height: 100%;
      width: 100%;
      border: none;
    }
    text-align: center;
    line-height: 40px;
    height: 40px;
  }
}

.dropdown {
  height: 100%;
  width: 100%;
}
v-overflow-btn {
  height: 100%;
  width: 100%;
}

v-row {
  height: 100%;
}
v-col {
  height: 100%;
}
</style>