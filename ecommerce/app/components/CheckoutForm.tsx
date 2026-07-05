interface Props {
    loading: boolean;
    onSubmit(): void;
}

export default function CheckoutForm({
    loading,
    onSubmit
}: Props) {

    return (

        <button
            disabled={loading}
            onClick={onSubmit}
        >

            {loading
                ? "Processing..."
                : "Place Order"}

        </button>

    );

}