import io.sentry.Sentry;
import io.sentry.SentryClient;
import io.sentry.SentryClientFactory;
import io.sentry.context.Context;
import io.sentry.event.BreadcrumbBuilder;
import io.sentry.event.UserBuilder;

public class MyClass {
    private static SentryClient sentry;

    public static void main(String... args) {
        // DSN direkt setzen
        Sentry.init("https://a1e3c23fc03643b0bffee804f5aa611b@glitchtip-stg.puzzle.ch/7");

        // Eigene Instanz holen
        sentry = SentryClientFactory.sentryClient();

        MyClass myClass = new MyClass();
        myClass.logWithStaticAPI();
        myClass.logWithInstanceAPI();
    }

    void unsafeMethod() {
        throw new UnsupportedOperationException("You shouldn't call this!");
    }

    void logWithStaticAPI() {
        Sentry.getContext().recordBreadcrumb(
            new BreadcrumbBuilder().setMessage("User made an action").build()
        );
        Sentry.getContext().setUser(
            new UserBuilder().setEmail("hello@sentry.io").build()
        );
        Sentry.getContext().addExtra("extra", "thing");
        Sentry.getContext().addTag("tagName", "tagValue");

        Sentry.capture("This is a test");

        try {
            unsafeMethod();
        } catch (Exception e) {
            Sentry.capture(e);
        }
    }

    void logWithInstanceAPI() {
        Context context = sentry.getContext();
        context.recordBreadcrumb(new BreadcrumbBuilder().setMessage("User made an action").build());
        context.setUser(new UserBuilder().setEmail("hello@sentry.io").build());

        sentry.sendMessage("This is a test");

        try {
            unsafeMethod();
        } catch (Exception e) {
            sentry.sendException(e);
        }
    }
}