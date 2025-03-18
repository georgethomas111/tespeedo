<template>
  <div>
    <div class="container">
      <h1>AI Response</h1>

      <div class="section">
        <div class="label">Model:</div>
        <div class="content">@cf/meta/llama-3.1-8b-instruct</div>
      </div>

      <form id="queryForm" @submit.prevent="handleSubmit">
        <input type="text" v-model="query" placeholder="Enter your question here..." />
        <input type="submit" value="Ask AI" />
      </form>

      <div class="section">
        <div class="label">Prompt:</div>
        <div class="content">{{ query }}</div>
      </div>

      <div class="section">
        <div class="label">Answer:</div>
        <div class="content">{{ answer || 'Waiting for your question...' }}</div>
      </div>

      <div class="section">
        <div class="label">Token Usage:</div>
        <div class="content">
          <strong>Prompt Tokens:</strong> {{ tokens.prompt || 0 }}<br>
          <strong>Completion Tokens:</strong> {{ tokens.completion || 0 }}<br>
          <strong>Total Tokens:</strong> {{ tokens.total || 0 }}
        </div>
      </div>

      <div class="error" v-if="error">
        <div class="label">Error:</div>
        <div class="content">{{ error }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      query: '',
      answer: '',
      tokens: { prompt: 0, completion: 0, total: 0 },
      error: ''
    };
  },
  methods: {
    async handleSubmit() {
      if (!this.query) return;

      // Clear previous error
      this.error = '';

      try {
        // Call the AI API
        const response = await fetch('/ask', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ query: this.query }),
        });

        const data = await response.json();

        // Update the answer and token usage sections
        if (data.response) {
          this.answer = data.response;
          this.tokens = {
            prompt: data.usage?.prompt_tokens || 0,
            completion: data.usage?.completion_tokens || 0,
            total: data.usage?.total_tokens || 0,
          };
        } else {
          this.answer = 'No response from AI';
        }
      } catch (error) {
        console.error("Error fetching AI response:", error);
        this.error = error.message || 'Error fetching AI response';
      }
    }
  }
};
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
  margin: 40px;
  background-color: #f9f9f9;
}
.container {
  max-width: 800px;
  margin: auto;
  background: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}
h1 {
  text-align: center;
  color: #333;
}
.section {
  margin-bottom: 20px;
}
.label {
  font-weight: bold;
  color: #555;
}
.content {
  padding: 10px;
  background: #f1f1f1;
  border-radius: 5px;
  white-space: pre-wrap;
}
form {
  margin-bottom: 20px;
}
input[type="text"] {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}
input[type="submit"] {
  padding: 10px 20px;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
input[type="submit"]:hover {
  background-color: #0056b3;
}
.error {
  background-color: #ffebee;
  color: #c62828;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 20px;
}
</style>

