<template>
  <section class="about-page">
    <h1>About Le'Blanc</h1>
    <div class="content">
      <p>
        Le'Blanc is not just a drink shop, but a <em>"personalized beverage space"</em>
        based on each customer's taste, health condition, mood, and purpose of use.
        At Le'Blanc, we carefully record and remember your individual taste profile,
        so each visit is a heart-warming experience.
        Combined with a calm, refined space and thoughtful service,
        we aim to turn every cup into a meaningful moment of comfort, clarity, or inspiration.
      </p>
      <p>
        We envision Le'Blanc as a place where every guest feels genuinely seen and understood.
        We aim to lead in personalized beverages by turning a simple drink into a thoughtful,
        tailored experience, so that when people think of <em>"a drink made just for me"</em>
        they naturally think of us.
      </p>
      <p>
        Our mission is to design drinks around each guest's taste, wellbeing, and mood,
        and to serve them with sincere, attentive care. We commit to listening,
        remembering preferences, and refining our menu and space around real needs,
        so every visit to Le'Blanc feels effortless, personal, and truly meaningful.
      </p>
    </div>
  </section>
</template>

<style scoped>
.about-page {
  max-width: min(1200px, 98vw);
  margin: 0 auto;
  padding: 48px 0 64px;
  display: grid;
  gap: 18px;
  color: var(--ink);
}

h1 {
  margin: 0;
  font-size: clamp(2rem, 4vw, 2.8rem);
  letter-spacing: 0.02em;
  font-weight: 700;
  font-style: italic;
  font-family: 'Georgia', 'Times New Roman', serif;
}

.content {
  display: grid;
  gap: 12px;
  font-size: 1.05rem;
  line-height: 1.7;
}

.content p {
  transition: transform 0.2s ease, color 0.2s ease;
}

.content p:hover {
  transform: translateY(-4px);
}

.content em {
  font-style: normal;
  font-weight: 800;
  transition: color 0.2s ease;
}

.content p:hover em {
  color: var(--tan);
}
</style>
