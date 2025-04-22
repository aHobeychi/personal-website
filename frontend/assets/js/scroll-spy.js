// Wrap entire script in an IIFE to prevent global variable redeclaration
(function() {
  // Define constants to avoid magic numbers
  const INIT_DELAY = 100;
  const INIT_DELAY_SHORT = 50;
  const INIT_DELAY_LONG = 300;
  const TOP_OFFSET = 120;
  const HEADING_PROXIMITY = 200;

  document.addEventListener('DOMContentLoaded', function () {
    // Initialize scrollspy both on page load and when content changes (HTMX)
    setTimeout(initScrollSpy, INIT_DELAY);

    // Also listen for HTMX content swap events to reinitialize when navigating between blogs
    document.body.addEventListener('htmx:afterSwap', function (event) {
      // Only reinitialize if the swap target is the main content
      if (event.detail.target.id === 'content-section') {
        setTimeout(initScrollSpy, INIT_DELAY);
      }
    });

    // Specifically listen for the TOC being loaded
    document.body.addEventListener('htmx:afterSettle', function (event) {
      // If the swap target is the sidebar container where TOC is loaded
      if (event.detail.target.id === 'variable-sidebar-container') {
        setTimeout(initScrollSpy, INIT_DELAY_SHORT);
      }
    });
  });

  // Clean up scroll spy resources
  function cleanupScrollSpy() {
    if (window.scrollSpyActive) {
      if (window.scrollObserver) {
        window.scrollObserver.disconnect();
        window.scrollObserver = null;
      }
      window.removeEventListener('resize', window.scrollSpyResizeHandler);
      window.removeEventListener('scroll', window.scrollSpyScrollHandler);
      window.scrollSpyActive = false;
    }
  }

  // Moved outside to be accessible from both event handlers
  function initScrollSpy() {
    try {
      // Only run if we have both a table of contents and article content
      const tocLinks = document.querySelectorAll('.blog-toc .toc-list a');
      const article = document.querySelector('article');
      if (!tocLinks.length || !article) {
        return;
      }
      // Create a map of headings used in the TOC
      const tocHeadings = {};
      tocLinks.forEach(link => {
        const href = link.getAttribute('href');
        if (href && href.startsWith('#')) {
          tocHeadings[href.substring(1)] = true;
        }
      });

      // Only get headings that are actually in the TOC
      const headings = Array.from(article.querySelectorAll('[id]')).filter(el =>
        tocHeadings[el.id]
      );

      if (!headings.length) {
        return;
      }

      // Clean up any previous scroll spy setup
      cleanupScrollSpy();

      // Track last active heading to prevent frequent changes
      let lastActiveHeading = null;
      // Create an array to store heading positions for better reference
      const headingPositions = [];

      // Pre-calculate initial heading positions and store them
      function calculateHeadingPositions() {
        headingPositions.length = 0; // Clear array
        headings.forEach(heading => {
          headingPositions.push({
            id: heading.id,
            top: heading.getBoundingClientRect().top + window.scrollY
          });
        });
        // Sort by position top to bottom
        headingPositions.sort((a, b) => a.top - b.top);
      }

      // Calculate initial positions
      calculateHeadingPositions();

      // Set a function to check which heading is in view
      function setActiveHeading() {
        const scrollPosition = window.scrollY;
        let currentHeadingId = null;

        // Find the current heading based on scroll position
        for (let i = headingPositions.length - 1; i >= 0; i--) {
          // If we've scrolled past this heading (with offset)
          if (scrollPosition + TOP_OFFSET >= headingPositions[i].top) {
            currentHeadingId = headingPositions[i].id;
            break;
          }
        }

        // Special case for the first heading when near the top of the page
        if (!currentHeadingId && headingPositions.length > 0 &&
          Math.abs(scrollPosition - headingPositions[0].top) < HEADING_PROXIMITY) {
          currentHeadingId = headingPositions[0].id;
        }

        // If we found a heading ID, update the active state
        if (currentHeadingId) {
          // Find the actual heading element
          const currentHeading = document.getElementById(currentHeadingId);

          // Only change if it's different from the last active heading
          if (!lastActiveHeading || currentHeadingId !== lastActiveHeading.id) {
            // Remove active class from all TOC links
            tocLinks.forEach(link => {
              link.classList.remove('active');
            });

            // Add active class to the current heading's TOC link
            const activeLink = document.querySelector(`.blog-toc .toc-list a[href="#${currentHeadingId}"]`);

            if (activeLink) {
              activeLink.classList.add('active');
              // Remember this heading as the last active one
              lastActiveHeading = currentHeading;
            }
          }
        }
      }

      // Implement Intersection Observer for better performance
      const observerOptions = {
        rootMargin: `-${TOP_OFFSET}px 0px -70% 0px`,
        threshold: [0, 0.1, 0.5, 1]
      };

      window.scrollObserver = new IntersectionObserver((entries) => {
        // Only process if we're not already animating
        if (!window.scrollSpyAnimating) {
          window.scrollSpyAnimating = true;
          window.requestAnimationFrame(() => {
            // If any heading is intersecting, we should recalculate
            if (entries.some(entry => entry.isIntersecting)) {
              setActiveHeading();
            }
            window.scrollSpyAnimating = false;
          });
        }
      }, observerOptions);

      // Observe all headings
      headings.forEach(heading => {
        window.scrollObserver.observe(heading);
      });

      // Store the resize handler so we can remove it later if needed
      window.scrollSpyResizeHandler = function () {
        calculateHeadingPositions();
        setActiveHeading();
      };

      // Add resize listener with passive flag for better performance
      window.addEventListener('resize', window.scrollSpyResizeHandler, { passive: true });

      // Mark that scroll spy is active
      window.scrollSpyActive = true;

      // Update heading positions when images or other resources finish loading
      window.addEventListener('load', function () {
        calculateHeadingPositions();
        setActiveHeading();
      });

      // Run once on load with a slight delay to ensure content is fully rendered
      setTimeout(function () {
        calculateHeadingPositions();
        setActiveHeading();
      }, INIT_DELAY_LONG);

      // Add a fallback scroll listener for browsers that might not fully support IntersectionObserver
      // This is in addition to the observer and helps ensure smooth operation
      window.scrollSpyScrollHandler = function () {
        if (!window.scrollSpyAnimating) {
          window.scrollSpyAnimating = true;
          window.requestAnimationFrame(function () {
            setActiveHeading();
            window.scrollSpyAnimating = false;
          });
        }
      };

      // Add scroll event with passive flag for performance
      window.addEventListener('scroll', window.scrollSpyScrollHandler, { passive: true });
    } catch (error) {
      console.error('Error initializing scroll spy:', error);
    }
  }
})();
