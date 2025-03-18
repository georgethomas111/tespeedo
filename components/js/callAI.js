import axios from 'axios';

const callAI = {
  async process(file) {
    try {
      const base64Image = await toBase64(file);

      const requestData = {
        model: "llava",
        prompt: "Describe this image in detail.",
        images: [base64Image],
      };

      const response = await axios.post("http://localhost:11434/api/generate", requestData, {
        headers: { "Content-Type": "application/json" },
      });

      console.log("AI Response:", response.data);
      alert(response.data.response); // Show response in an alert
    } catch (error) {
      console.error("Error processing image:", error);
      alert("Failed to process image. Please try again.");
    }
  },
};

// Convert file to Base64
function toBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
      // Extract base64 content without the metadata prefix (e.g., "data:image/png;base64,")
      const base64String = reader.result.split(',')[1];
      resolve(base64String);
    };
    reader.onerror = (error) => reject(error);
  });
}

export default callAI;

