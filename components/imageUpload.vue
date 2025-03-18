<template>
  <div class="image-upload">
    <input type="file" @change="onFileChange" accept="image/*" />

    <p v-if="errorMessage" class="error">{{ errorMessage }}</p>

    <div v-if="imageUrl" class="image-preview">
      <img :src="imageUrl" alt="Image preview" />
    </div>

    <button v-if="imageUrl" @click="clearImage">Clear Image</button>
  </div>
</template>

<script>
import callAI from './js/callAI.js'

export default {
  name: "ImageUpload",
  
  data() {
    return {
      imageUrl: null,      // URL of the uploaded image to display preview
      errorMessage: null,  // Error message if there are issues with the file
    };
  },

  methods: {
    // Handle file input change event
    onFileChange(event) {
      const file = event.target.files[0];
      this.errorMessage = null;  // Clear any previous error messages

      // Validate file type and size
      if (file && file.type.startsWith("image/")) {
        if (file.size > 2 * 1024 * 1024) {  // Limit size to 2MB
          this.errorMessage = "Image size should be less than 2MB.";
          return;
        }

        // Create an image preview
        this.imageUrl = URL.createObjectURL(file);
	callAI.process(file)
      } else {
        this.errorMessage = "Please upload a valid image file.";
      }
    },

    // Clear the uploaded image
    clearImage() {
      this.imageUrl = null;
      this.$refs.fileInput.value = null;  // Reset the file input
    },
  },
};
</script>

<style scoped>
.image-upload {
  max-width: 60%;
  margin: 0 auto;
  text-align: left;
}

input[type="file"] {
  margin: 10px 0;
}

.image-preview img {
  max-width: 100%;
  height: auto;
  margin-top: 10px;
}

.error {
  color: red;
}
</style>
